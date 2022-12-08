import Ajv from "ajv";

export const validateBody = (schema: any) => {
  const ajv = new Ajv();

  const validate: any = ajv.compile(schema);

  return (req: any, res: any, next: any) => {
    const valid = validate(req.body);

    if (valid) {
      return next();
    } else {
      const error = validate.errors[0];

      return res.status(400).json({
        error: {
          message: `"${error.instancePath.substring(1)}" ${error.message}`,
        },
      });
    }
  };
};
