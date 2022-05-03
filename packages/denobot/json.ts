export async function readJson(filePath: string): Promise<unknown> {
  try {
    const jsonString = await Deno.readTextFile(filePath);
    return JSON.parse(jsonString);
  } catch (err) {
    err.message = `${filePath}: ${err.message}`;
    throw err;
  }
}

/** Reads a JSON file and then parses it into an object */
export function readJsonSync(filePath: string): unknown {
  try {
    const jsonString = Deno.readTextFileSync(filePath);
    return JSON.parse(jsonString);
  } catch (err) {
    err.message = `${filePath}: ${err.message}`;
    throw err;
  }
}
