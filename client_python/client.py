import sys
import grpc
import greeting_pb2
import greeting_pb2_grpc

def run_greeter():
    with grpc.insecure_channel('localhost:50051') as channel:
        greeter_stub = greeting_pb2_grpc.GreeterStub(channel)

        # Call the greet method
        response = greeter_stub.greet(greeting_pb2.ClientInput(name='John', greeting='Yo'))
        print("Greeter client received following from server: " + response.message)


def run_book():
    with grpc.insecure_channel('localhost:50051') as channel:
        book_stub = greeting_pb2_grpc.BookStoreStub(channel)

        # Call the first method of the BookStore service
        book_request = greeting_pb2.BookSearch(name='Random Book', author='Cool Guy', genre='Nice')
        book_response = book_stub.first(book_request)
        print("BookStore client received the following book from server: " + str(book_response))


def run_movie():
    with grpc.insecure_channel('localhost:50051') as channel:
        movie_stub = greeting_pb2_grpc.MovieClubStub(channel)

        # Call the first method of the movieStore service
        movie_request = greeting_pb2.MovieSearch(name='Random movie', director='Cool Guy', genre='Nice')
        movie_response = movie_stub.first(movie_request)
        print("Movie Club client received the following book from server: " + str(movie_response))

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Usage: python client.py <function_name>")
        sys.exit(1)

    function_name = sys.argv[1]

    if function_name == 'run_greeter':
        run_greeter()
    elif function_name == 'run_book':
        run_book()
    elif function_name == 'run_movie':
        run_movie()
    else:
        print(f"Unknown function: {function_name}")
        sys.exit(1)