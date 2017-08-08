var PROTO_PATH_API = __dirname + '/../../api/api.proto';
var PROTO_PATH = __dirname + '/../../phonebook.proto';
var grpc = require('grpc');
var path = require('path');
console.log(PROTO_PATH);
console.log("Attempting to load protopath");
var protoDesc = grpc.load(PROTO_PATH);
console.log(protoDesc);
var apiDesc = grpc.load(path.resolve(PROTO_PATH_API));
var client = new phonebook.PhoneBook('localhost:50051', grpc.credentials.Insecure());

var inquirer = require('inquirer');

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
    console.log("other answers");
    console.log(answers);
    if (!answers.createConfirm) {
      promptCreateContact(); 
    } else {
      // Create the contact in the gRPC service.
      // We'll call the gRPC CreateContact endpoint from here.
      // Redirect to the root menu
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
      return []
    },
    type: 'rawlist'
  }];

  inquirer.prompt(listQuestions).then(function(answers) {
    console.log("In the listContact method.");
    // We probably will want to output the user information here and then ask
    // the user what the hell they want to do about it?
  
    // Once we setup another prompt for the inquirer 
    console.log(answers);
    ask();
  });
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
    console.log("first answers: " + JSON.stringify(answers));
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
