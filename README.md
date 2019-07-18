# golang_rhel_os_check
A simple module to fetch the RedHat release version at runtime.

$cat /etc/redhat-release

Red Hat Enterprise Linux Server release 6.5 (Santiago)

I am reading in the 40th character which in the above examples reads in "6".

I am reading in a one-byte character only as opposed to the whole file contents in order to reduce space complexity.
