var PROTO_PATH = __dirname + '/../../api/api.proto';
var grpc = require('grpc');
var path = require('path');
var inquirer = require('inquirer');

var api = grpc.load(path.resolve(PROTO_PATH)).api;
var client = new api.PhoneBook('localhost:50051', grpc.credentials.createInsecure());
var PhoneType = api.PhoneNumber.PhoneType;

function promptCreateContact() {
  var createQuestions = [
    {
      name: 'createName',
      type: 'input',
      message: "What is this contact's name?",
      validate: function(input) {
        return input.trim().length > 0 || 'Please enter a name';
      },
    },
    {
      name: 'createEmail',
      type: 'input',
      message: "What is this contact's email address?",
      when: function(answers) {
        return answers.createName;
      },
    },
    {
      name: 'createPhoneNumbers',
      type: 'checkbox',
      message: "What type of numbers would you like to add?",
      choices: ['home', 'work', 'mobile']
    },
    {
      name: 'createMobilePhoneNumber',
      type: 'input',
      message: "What is their mobile phone number?",
      when: function(answers) {
        return answers.createPhoneNumbers.indexOf("mobile") >= 0;
      }
    },
    {
      name: 'createHomeNumber',
      type: 'input',
      message: "What is their home phone number?",
      when: function(answers) {
        return answers.createPhoneNumbers.indexOf("home") >= 0;
      }
    },
    {
      name: 'createWorkNumber',
      type: 'input',
      message: "What is their work phone number?",
      when: function(answers) {
        return answers.createPhoneNumbers.indexOf("work") >= 0;
      }
    },
    {
      type: 'confirm',
      name: 'createConfirm',
      message: 'Does everything look right to you?',
    default: true
    }
  ];
  inquirer.prompt(createQuestions).then(function(answers) {
    if (!answers.createConfirm) {
      promptCreateContact(); 
    } else {
      // Create the contact in the gRPC service.
      // We'll call the gRPC CreateContact endpoint from here.
      // Redirect to the root menu
      var request = new api.CreateContactReq();
      var contact = new api.Contact();
      var phoneArr = [];

      if (answers.createHomeNumber) {
        var homeNum = new api.PhoneNumber();
        homeNum.type = PhoneType.HOME;
        homeNum.number = answers.createHomeNumber;
        phoneArr.push(homeNum)
      }

      if (answers.createWorkNumber) {
        var workNum = new api.PhoneNumber();
        workNum.type = PhoneType.WORK;
        workNum.number = answers.createWorkNumber;
        phoneArr.push(workNum)
      }

      if (answers.createMobileNumber) {
        var mobileNum = new api.PhoneNumber();
        mobileNum.type = PhoneType.MOBILE;
        mobileNum.number = answers.createMobileNumber;
        phoneArr.push(mobileNum)
      }

      request.name = answers.createName;
      request.email = answers.createEmail;
      request.phone_numbers = phoneArr;
      client.createContact(request, function(err, response) {
        if (err) {
          console.log("Error mentioned!");
          console.log(err);
          return;
        }
      });
      ask(); 
    }
  });
}

function promptListContacts() {
  // Maybe list contacts should be `view` contacts?
  // And then have a separate action for editing?
  // And another for deletion.
  var listQuestions = [{
    name: 'contacts',
    message: 'Which contact do you want to view?',
    choices: function(answers) {
      // Call the ListContacts gRPC endpoint here, map it to an array to return.
      var done = this.async();
      console.log("Get the choices here.");
      client.listContacts({}, function(err, response) {
        if (err) {
          // We want to return the error here.
          console.log("Error returned!");
          console.log(err);
          return;
        }

        var choices = response.contacts.map(function(contact) {
          return {
            name: contact.name,
            value: contact.name,
          };
        });

        done(null, choices);
      });
    },
    type: 'rawlist'
  }];

  inquirer.prompt(listQuestions).then(function(answers) {
    // We probably will want to output the user information here and then ask
    // the user what the hell they want to do about it?

    // Once we setup another prompt for the inquirer 
    ask();
  });
}

function promptShowContact() {

}

function ask() {
  var promptQuestions = [
    {
      name: 'action',
      type: 'rawlist',
      message: 'What would you like to do?',
      choices: [
        {
          name: 'Add new contact',
          value: 'create'
        },
        {
          name: 'List Contacts',
          value: 'list'
        }
      ]
    },
  ];
  inquirer.prompt(promptQuestions).then(function(answers) {
    if (answers.action == 'create') {
      promptCreateContact(); 
    } else if (answers.action = 'list') {
      // List all this stuff 
      promptListContacts();
    } else {
      // Re-run the ask prompt 
      ask()
    }
  });
}

ask();
