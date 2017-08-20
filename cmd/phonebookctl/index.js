var PROTO_PATH = __dirname + '/../../api/api.proto';
var grpc = require('grpc');
var program = require('commander');
var path = require('path');
var api = grpc.load(path.resolve(PROTO_PATH)).api;
var client = new api.PhoneBook('localhost:50051', grpc.credentials.createInsecure());

var PhoneType = api.PhoneNumber.PhoneType;

const VERSION = '0.0.1';

function phoneNumberArr(val, numbers) {
  numbers.push(val);
  return numbers;
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
      });

    }

    if (options.work) {
      options.work.forEach(function(num) {
        var workNum = new api.PhoneNumber();
        workNum.type = PhoneType.WORK;
        workNum.number = num;
        phoneArr.push(workNum);
      });

    }

    if (options.mobile) {
      options.mobile.forEach(function(num) {
        var mobileNum = new api.PhoneNumber();
        mobileNum.type = PhoneType.MOBILE;
        mobileNum.number = num;
        phoneArr.push(mobileNum);
      });
    }

    request.name = name;
    request.email = options.email || '';
    request.phone_numbers = phoneArr;
    client.createContact(request, function(err, response) {
      if (err) {
        console.log("There's an error.");
        console.log(err);
        return;
      }

      console.log(response);
    });
  });


program.parse(process.argv);
