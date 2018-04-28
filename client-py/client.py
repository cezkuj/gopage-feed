import gopage_pb2
import gopage_pb2_grpc
import grpc

channel = grpc.insecure_channel('localhost:8002')
stub = gopage_pb2_grpc.GopageStub(channel)
print(stub.get1(gopage_pb2.GetParams()).value)
