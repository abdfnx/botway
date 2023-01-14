export const EmailTemplate = (
  username: string,
  tokenId: string,
  webURL: any,
  path: string
) => {
  const shared = `color: #000; font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif; font-size:`;

  let message, protocol;

  if (path == "verify-email") {
    message = "to confirm your email";
  } else if (path == "forget-password") {
    message = "to reset your password";
  }

  if (webURL.includes("localhost")) {
    protocol = "http";
  } else {
    protocol = "https";
  }

  return `<div>
      <table
        width="100%"
        border="0"
        cellspacing="0"
        cellpadding="0"
        style="width: 100% !important"
      >
        <tbody>
          <tr>
            <td align="center">
              <table
                style="border: 1px solid #eaeaea; border-radius: 5px; margin: 40px 0;"
                width="600"
                border="0"
                cellspacing="0"
                cellpadding="40"
              >
                <tbody>
                  <tr>
                    <td align="center">
                      <div style="text-align: left; width: 465px">
                        <table
                          width="100%"
                          border="0"
                          cellspacing="0"
                          cellpadding="0"
                          style="width: 100% !important"
                        >
                          <tbody>
                            <tr>
                              <td align="center">
                                <img
                                  src="https://cdn-botway.deno.dev/icon.png"
                                  width="40px"
                                  alt="Botway"
                                />
                                <h1 style="${shared} 24px; font-weight: normal; margin: 30px 0; padding: 0;">
                                  Here's your Botway Verification URL
                                </h1>
                              </td>
                            </tr>
                          </tbody>
                        </table>
                        <p style="${shared} 14px; line-height: 24px;">
                          Hi <strong>${username}</strong>
                        </p>
                        <p style="${shared} 14px; line-height: 24px;">
                          Please follow <a href="${protocol}://${webURL}/${path}/${tokenId}">this link</a> ${message}.
                        </p>
                        <br />

                        <br />
                        <hr style="border: none; border-top: 1px solid #eaeaea; margin: 26px 0; width: 100%;" />
                        <p style="color: #666666; font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif; font-size: 12px; line-height: 24px;">
                          &copy; Botway. All rights reserved.
                        </p>
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </td>
          </tr>
        </tbody>
      </table>
    </div>`;
};
