// @deno-types="npm:@types/express"
import express from "npm:express";
import _ from "npm:pug";

const app = express();

app.use(express.static("cdn"));

app.set("views", new URL(".", import.meta.url).pathname);
app.set("view engine", "pug");

app.use(express.json());

app.use(express.urlencoded({ extended: false }));

app.get("/", (req, res) => {
  res.render("layout", {
    message: JSON.stringify(
      { status: 200, message: "welcome to botway cdn 📦" },
      null,
      " "
    ).toString(),
  });
});

app.listen(8000);
