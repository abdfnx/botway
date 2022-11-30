export const ValidateProps = {
  user: {
    username: { type: "string", minLength: 4, maxLength: 20 },
    name: { type: "string", minLength: 1, maxLength: 50 },
    password: { type: "string", minLength: 8 },
    email: { type: "string", minLength: 1 },
    isAdmin: { type: "boolean" },
  },
  project: {
    name: { type: "string", minLength: 1, maxLength: 350 },
    platform: { type: "string", minLength: 5, maxLength: 8 },
    lang: { type: "string", minLength: 1, maxLength: 50 },
    package_manager: { type: "string", minLength: 1, maxLength: 50 },
    host_service: { type: "string", minLength: 1, maxLength: 50 },
  },
};
