var PROTO_PATH = __dirname + '/../../api/api.proto';
var grpc = require('grpc');
var program = require('commander');
var path = require('path');
var api = grpc.load(path.resolve(PROTO_PATH)).api;
var client = new api.PhoneBook('localhost:50051', grpc.credentials.createInsecure());
require('console.table');

var PhoneType = api.PhoneNumber.PhoneType;

const VERSION = '0.0.1';

function phoneNumberArr(val, numbers) {
  numbers.push(val);
  return numbers;
}

function printHeader() {
  console.log("Name\t\tEmail\t\tHome\t\tMobile\t\tWork\t\t");
}

function formatContacts(contacts) {
  return contacts.map(function(contact) {
    newContact = contact;
    newContact.home = contact.home && contact.home.number || '';
    newContact.mobile = contact.mobile && contact.mobile.number || '';
    newContact.work  = contact.work && contact.work.number || '';
    delete newContact.phone_numbers;

    return newContact;
  });
}

function printContacts(contacts) {
  console.table(contacts);
}

function printContact(contact) {
  console.log(tabItem(contact.name) + tabItem(contact.email));
}

function tabItem(value) {
  return value + "\t\t";
}

program
  .version(VERSION)

program
  .command('create <name>')
  .option('-n, --name <name>', "Contact's name")
  .option('-e, --email <email>', "Contact's email")
  .option('-h, --home <home>', "Contact's home number", phoneNumberArr, [])
  .option('-m, --mobile <mobile>', "Contact's mobile number", phoneNumberArr, [])
  .option('-w, --work <work>', "Contact's work number", phoneNumberArr, [])
  .action(function(name, options) {
    var request = new api.CreateContactReq();
    var phoneArr = [];

    if (options.home) {
      options.home.forEach(function(num) {
        var homeNum = new api.PhoneNumber();
        homeNum.type = PhoneType.HOME;
        homeNum.number = num;
        phoneArr.push(homeNum);
        request.home = num;
      });
    }

    if (options.work) {
      options.work.forEach(function(num) {
        var workNum = new api.PhoneNumber();
        workNum.type = PhoneType.WORK;
        workNum.number = num;
        phoneArr.push(workNum);
        request.work = num;
      });
    }

    if (options.mobile) {
      options.mobile.forEach(function(num) {
        var mobileNum = new api.PhoneNumber();
        mobileNum.type = PhoneType.MOBILE;
        mobileNum.number = num;
        phoneArr.push(mobileNum);
        request.mobile = num;
      });
    }

    request.name = name;
    request.email = options.email || '';
    request.phone_numbers = phoneArr;
    request.home = request.home || '';
    request.mobile = request.mobile || '';
    request.work = request.work || '';

    client.createContact(request, function(err, response) {
      if (err) {
        console.log("There's an error.");
        console.log(err);
        return;
      }

      console.log(response);
    });
  });

program
  .command('list')
  .action(function() {
    var request = new api.ListContactsReq();
    client.listContacts(request, function(err, response) {
      if (err) {
        console.log("There's an error");
        console.log(err);
        return;
      }

      printContacts(formatContacts(response.contacts));
    });
  });

program.parse(process.argv);
