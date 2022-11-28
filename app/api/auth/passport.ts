import { findUserForAuth, findUserWithEmailAndPassword } from "api/db";
import passport from "passport";
import { Strategy as LocalStrategy } from "passport-local";
import { getMongoDb } from "../mongodb";

passport.serializeUser((user: any, done: any) => {
  done(null, user._id);
});

passport.deserializeUser((req: any, id: any, done: any) => {
  getMongoDb().then((db) => {
    findUserForAuth(db, id).then(
      (user) => done(null, user),
      (err) => done(err)
    );
  });
});

passport.use(
  new LocalStrategy(
    { usernameField: "email", passReqToCallback: true },
    async (req: any, email: any, password: any, done: any) => {
      const db = await getMongoDb();
      const user = await findUserWithEmailAndPassword(db, email, password);

      if (user) done(null, user);
      else done(null, false, { message: "Email or password is incorrect" });
    }
  )
);

export default passport;
