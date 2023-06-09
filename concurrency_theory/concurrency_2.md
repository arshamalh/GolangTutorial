Channels aren't resources like files or sockets, you don't need to close them to free them.
Closing a channels is just a signal for the receiver that it doesn't have any other data.