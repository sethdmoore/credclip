# Introduction

Grabs AWS credentials from your ~/.aws/credentials file and copies the following string to your clipboard.  

```
unset HISTFILE; export AWS_ACCESS_KEY_ID="AKIAXXXXXXXXXXXXXXXX"; export AWS_SECRET_ACCESS_KEY="arBOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX";
```

Good for pasting credentials into a remote session. That's about it.  

# Usage

By default, credclip will grab the "default" section of your AWS credentials.  
If you want to select a different section, EG:  

```
[default]
aws_access_key_id = AKIAXXXXXXXXXXXXXXXX
aws_secret_access_key = arBOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

[production]
aws_access_key_id = AKIAXXXXXXXXXXXXXXXX
aws_secret_access_key = RoNtXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

Select the section as the first argument, EG:  
```
$ credclip production
```
