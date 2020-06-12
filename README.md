# random-file-generator

This is a lightweight and portable utility for generating a number of files of specified size. These files are often used for testing storage and data transfer solutions. 

The content of these files consists of randomized bytes, not NULL content like often found in big test files on the Internet. This is important for situations where data deduplication is in play.

## Usage
You must specify at least the number of files with the -Count flag.
```
  -Count int
        [Required] Number of files to generate

  -Dir string
        Target directory for files. 
        Will be created if not exists. 
        Defaults to the current working directory.

  -Size string
        Filesizes in kilobytes. 
        Single value or comma seperated for random distribution
        Default is 32.
```

### Example
Generate 5000 files in a test subdirectory of the current directory. Size to be evenly distributed as 32,64,128 and 256 kilobytes

`random-file-generator -Count 5000 -Dir test -Size 32,64,128,256`