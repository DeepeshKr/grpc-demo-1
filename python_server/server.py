import signal

import grpc
from concurrent import futures
import signal
import greeting_pb2
import greeting_pb2_grpc

class Greeter(greeting_pb2_grpc.GreeterServicer):
    def greet(self, request, context):
        print("Greeting request " + str(request))
        return greeting_pb2.ServerOutput(message='{0} {1}!'.format(request.greeting, request.name))

class BookStore(greeting_pb2_grpc.BookStoreServicer):
    def first(self, request, context):
        print("Book request " + str(request))
        # Implement logic to search for the book and return the result
        book = greeting_pb2.Book(name=request.name, author=request.author, price=20)
        return book

class MovieClub(greeting_pb2_grpc.MovieClubServicer):
    def first(self, request, context):
        print("Movie request " + str(request))
        # Implement logic to search for the movie and return the result
        movie = greeting_pb2.Movie(name=request.name, director=request.director, rating=6.1)
        return movie

def server():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=2))
    greeting_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    greeting_pb2_grpc.add_BookStoreServicer_to_server(BookStore(), server)
    greeting_pb2_grpc.add_MovieClubServicer_to_server(MovieClub(), server)

    signal.signal(signal.SIGINT, lambda sig, frame: server.stop(0))

    server.add_insecure_port('[::]:50051')
    print("gRPC starting")
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    server()


def signal_handler(sig, frame):
    print('Received signal, shutting down gracefully...')
    server.stop(0)

signal.signal(signal.SIGINT, signal_handler)