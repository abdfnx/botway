import { createWriteStream } from "fs";
import * as fs from "fs/promises";
import fetch from "node-fetch";
import { pipeline } from "stream/promises";
import StreamZip from "node-stream-zip";
import { ARCH_MAPPING, CONFIG, PLATFORM_MAPPING } from "./config.js";
import path from "path";

async function install() {
  const packageJson = await fs.readFile("package.json").then(JSON.parse);
  let version = packageJson.version;

  if (typeof version !== "string") {
    throw new Error("Missing version in package.json");
  }

  if (version[0] === "v") version = version.slice(1);

  // Fetch Static Config
  let { name: binName, url: url } = CONFIG;

  url = url.replace(/{{arch}}/g, ARCH_MAPPING[process.arch]);
  url = url.replace(/{{platform}}/g, PLATFORM_MAPPING[process.platform]);
  url = url.replace(/{{version}}/g, version);
  url = url.replace(/{{bin_name}}/g, binName);

  const folder = (old) => {
    let f =
      binName +
      "_" +
      PLATFORM_MAPPING[process.platform] +
      "_" +
      "v" +
      version +
      "_" +
      ARCH_MAPPING[process.arch];

    if (old == "yes") {
      return path.join(f, "bin");
    } else if (old == "no") {
      return "bin";
    } else {
      return f;
    }
  };

  const response = await fetch(url);

  console.log(response);

  if (!response.ok) {
    throw new Error("Failed fetching the binary: " + response.statusText);
  }

  const zipFile = "botway.zip";

  await pipeline(response.body, createWriteStream(zipFile));
  const zip = new StreamZip.async({ file: zipFile });

  const count = await zip.extract(null, ".");

  console.log(`Extracted ${count} entries`);

  await zip.close();

  await fs.rename(folder("yes"), folder("no"), function (err) {
    if (err) throw err;
  });

  await fs.rm(zipFile);
  await fs.rm(folder(), { recursive: true });
}

install()
  .then(async () => {
    process.exit(0);
  })
  .catch(async (err) => {
    console.error(err);
    process.exit(1);
  });
