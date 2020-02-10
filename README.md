An example of a situation when RPC is not processed as expected 

# How to run?
It requires Go modules.

```
git clone git@github.com:peti2001/test-micro-call.git
cd test-micro-call
cd serviceA
go run main.go
```
in another terminal start ServiceB
```
cd test-micro-call
cd serviceB
go run main.go
```

# Output
ServiceA
```
2020-02-10 22:21:21.265350 I | Transport [http] Listening on [::]:65409
2020-02-10 22:21:21.265437 I | Broker [http] Connected to [::]:65410
2020-02-10 22:21:21.265730 I | Registry [mdns] Registering node: serviceA-f32ac091-6d90-48e9-b3c6-0d414275ced3
2020-02-10 22:21:21.269132 I | Subscribing serviceA-f32ac091-6d90-48e9-b3c6-0d414275ced3 to topic: sayHello.topic
2020-02-10 22:21:23.417395 I | Consume message. MessageId: 1021023610000
2020-02-10 22:21:23.417415 I | Hello, Peter
DEBUG ctx &context.valueCtx{Context:(*context.emptyCtx)(0xc0000420b8), key:metadata.metaKey{}, val:metadata.Metadata{"Content-Type":"application/protobuf", "Micro-From-Service":"serviceB", "Micro-Id":"f55d1c0c-73a0-48d4-823d-9f31b6eccd3d", "Micro-Topic":"sayHello.topic"}}
2020-02-10 22:21:25.524018 I | Ack is sent to ServiceB. Status:
2020-02-10 22:21:25.525139 I | Ack is sent to ServiceB with new context. Status: ok
```
ServiceB
```
2020-02-10 22:21:21.308042 I | Transport [http] Listening on [::]:65413
2020-02-10 22:21:21.308150 I | Broker [http] Connected to [::]:65414
2020-02-10 22:21:21.308417 I | Registry [mdns] Registering node: serviceB-b58169f2-5afd-4f82-aa68-87a68a48c1d2
2020-02-10 22:21:23.415727 I | Ask ServiceA to do a long process by publishing a message. Waiting for ack so I can continue after that. MessageId: 1021023610000
2020-02-10 22:21:25.524491 I | Ack recieved of messageId 1021023610000
```
