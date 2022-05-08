#!/usr/bin/env node
import { execFileSync } from "child_process";
import { exit } from "process";

try {
  execFileSync("./botwaybin", process.argv.slice(2), {
    stdio: "inherit",
  });
} catch (e) {
  console.log(e);
  exit(1);
}
