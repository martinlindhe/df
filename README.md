# About

`df` prints the disk usage / bytes free of partitions.

Re-implementation of the `df` utility from coreutils,
written in Golang and targeting Windows.


# Installation

    go get -u github.com/martinlindhe/go-df/cmd/df


# TODO

* rework --human-readable to show appropriate units, now only gigabyte (G) is used


# Why make this?

Using windows commands

    $ fsutil volume diskfree c:
    The FSUTIL utility requires that you have administrative privileges.


    $ wmic /node:"%COMPUTERNAME%" LogicalDisk Where DriveType="3" Get DeviceID,FreeSpace|find /I "c:"
    Node - %COMPUTERNAME%
    ERROR:
    Description = The RPC server is unavailable.

none of suggestions at https://stackoverflow.com/questions/293780/free-space-in-a-cmd-shell works for me


## License

Under [MIT](LICENSE)