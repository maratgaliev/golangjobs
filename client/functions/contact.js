const aws = require("aws-sdk");

// Add your AWS credentials here
aws.config.update({
  accessKeyId: "",
  secretAccessKey: "",
  region: "us-west-2"
});

// Load AWS SES
const ses = new aws.SES({ apiVersion: "2010-12-01" });
// Add your email address here
const to = [""];
// Also add your email address as the sender
// Must belong to a verified SES account
const from = "";

exports.handler = function(event, context, callback) {
  const body = JSON.parse(event.body);

  ses.sendEmail(
    {
      Source: from,
      Destination: { ToAddresses: to },
      Message: {
        Subject: {
          Data: `Contact form submission`
        },
        Body: {
          Text: {
            Data: `
                  Name: ${body.name}
                  Email: ${body.email}
                  Message: ${body.message}
                `
          }
        }
      }
    },
    (err, data) => {
      if (err) {
        callback(null, {
          statusCode: 200,
          body: JSON.stringify({ status: "error" })
        });
      } else {
        callback(null, {
          statusCode: 200,
          body: JSON.stringify({ status: "success" })
        });
      }
    }
  );
};
