var path = require('path');
var PROTO_PATH = path.resolve(__dirname, '../helloworld/helloworld/helloworld.proto');

var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
var hello_proto = grpc.loadPackageDefinition(packageDefinition).helloworld;
const client = new hello_proto.Greeter('localhost:50051', grpc.credentials.createInsecure());

module.exports = client