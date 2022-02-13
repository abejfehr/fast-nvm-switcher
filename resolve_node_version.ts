import { dirname, join } from "https://deno.land/std/path/mod.ts";
import { existsSync } from "https://deno.land/std/fs/mod.ts";
import * as semver from "https://deno.land/x/semver/mod.ts";

// this program's whole purpose is to just output to stdout the path to the node version to the shell can put it on the path

const getNvmrcPath = () => {
  let pwd = Deno.cwd();

  let nvmrcPath = join(pwd, '.nvmrc');

  if (existsSync(nvmrcPath)) {
    return nvmrcPath;
  }

  while (pwd !== '/') {
    pwd = dirname(pwd); // goes up one

    const nvmrcPath = join(pwd, '.nvmrc');

    if (existsSync(nvmrcPath)) {
      return nvmrcPath;
    }
  }
}

const getNodePath = async (version: string) => {
  const versions: string[] = [];

  const isFuzzyVersion = !semver.valid(version, { loose: true })

  for await (const dirEntry of Deno.readDir(`${Deno.env.get('NVM_DIR')}/versions/node/`)) {
    if (dirEntry.isDirectory) {
      versions.push(dirEntry.name);
    }
  }

  const resolvedVersion = semver.maxSatisfying(versions, `${isFuzzyVersion ? '^' : ''}${version}`);

  if (!resolvedVersion) {
    console.error("Unable to find node version that matches, please run 'nvm install'");
    Deno.exit(1);
  }

  return `${Deno.env.get('NVM_DIR')}/versions/node/${resolvedVersion}/bin`;
}

const main = async () => {
  const nvmrcPath = getNvmrcPath();

  let version: string | undefined;

  if (nvmrcPath) {
    version = await Deno.readTextFile(nvmrcPath);
  }

  if (!version) {
    // Get the nvm "default" alias
    version = await Deno.readTextFile(`${Deno.env.get('NVM_DIR')}/alias/default`)
  }

  const nodePath = await getNodePath(version);

  console.log(nodePath)
}

main();
