const sgMail = require("@sendgrid/mail");

sgMail.setApiKey(process.env.SENDGRID_API_KEY);

export const sendMail = (email: any, subject: string, html: string) => {
  const msg = {
    to: email,
    from: process.env.EMAIL_FROM,
    subject,
    html,
  };

  sgMail
    .send(msg)
    .then(() => {
      console.log("Email sent");
    })
    .catch((error: any) => {
      console.error(error);
    });
};
