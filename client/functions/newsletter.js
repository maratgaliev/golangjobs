const Mailchimp = require("mailchimp-api-v3");

// Add your Mailchimp credentials here
const apiKey = "";
const listId = "";
const mailchimp = new Mailchimp(apiKey);

exports.handler = function(event, context, callback) {
  const body = JSON.parse(event.body);

  mailchimp
    .request({
      method: "POST",
      path: "/lists/" + listId + "/members",
      body: {
        email_address: body.email,
        // Set status to "subscribed" to disable double-opt-in
        status: "pending"
      }
    })
    .then(result => {
      callback(null, {
        statusCode: 200,
        body: JSON.stringify({ status: "success" })
      });
    })
    .catch(err => {
      callback(null, {
        statusCode: 200,
        body: JSON.stringify({ status: "error" })
      });
    });
};
