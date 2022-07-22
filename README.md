sls offline --stage dev


openssl dgst -sha256 -binary postConfirmation.zip | openssl enc -base64 | tr -d "\n" > postConfirmation.zip.base64sha256

aws lambda update-function-code --function-name postConfirmation --s3-bucket bitsports-lambda --s3-key postConfirmation.zip
