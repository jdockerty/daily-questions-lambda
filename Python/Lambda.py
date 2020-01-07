import json
import boto3
from datetime import datetime


def lambda_handler(event, context):
    ses_client = client_setup()
    source, dest, message = email_setup()
    ses_client.send_email(Source=source, Destination=dest, Message=message)
    current_time = datetime.now().strftime('%d-%m-%Y %H:%M:%S')
    print("Email sent at: ", current_time)
    return True


def client_setup():
    client = boto3.client('ses')
    return client
    
def email_setup():
    email_source = "[MY_EMAIL]"
    email_destination = {
        'ToAddresses': ['MY_EMAIL'], 
        'CcAddresses': [], 
        'BccAddresses': []
    }
    email_message_contents = { 'Body': 
        { 'Html': 
            {
            'Charset': 'UTF-8', 
            'Data': 
'<h3>Lambda in the Oregon Region</h3><h4>Did I do my best to...</h4> <ul> <li>place questions here</li> <li> etc... </li></ul>'
        },
            'Text': {
                'Charset': 'UTF-8',
                'Data': '',
            },
        },
        'Subject': {
            'Charset': 'UTF-8',
            'Data': 'Daily Questions',
        },
    }
    return email_source, email_destination, email_message_contents
