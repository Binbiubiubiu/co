import {
  dirname,
  fromFileUrl,
  resolve,
  extname,
  basename,
} from "https://deno.land/std@0.177.0/path/mod.ts";
// @deno-types="npm:@types/ejs@^3.1.2"
import * as ejs from "npm:ejs@^3.1.8";
import { Style } from "./types.ts";

const styles: Array<Style> = [
  { name: "Reset", open: 0, close: 0 },
  { name: "Bold", open: 1, close: 22, replace: '"\\x1b[22m\\x1b[1m"' },
  { name: "Dim", open: 2, close: 22, replace: '"\\x1b[22m\\x1b[2m"' },
  { name: "Italic", open: 3, close: 23 },
  { name: "Underline", open: 4, close: 24 },
  { name: "Inverse", open: 7, close: 27 },
  { name: "Hidden", open: 8, close: 28 },
  { name: "Strikethrough", open: 9, close: 29 },
  { name: "Black", open: 30, close: 39 },
  { name: "Red", open: 31, close: 39 },
  { name: "Green", open: 32, close: 39 },
  { name: "Yellow", open: 33, close: 39 },
  { name: "Blue", open: 34, close: 39 },
  { name: "Magenta", open: 35, close: 39 },
  { name: "Cyan", open: 36, close: 39 },
  { name: "White", open: 37, close: 39 },
  { name: "Gray", open: 90, close: 39 },
  { name: "BgBlack", open: 40, close: 49 },
  { name: "BgRed", open: 41, close: 49 },
  { name: "BgGreen", open: 42, close: 49 },
  { name: "BgYellow", open: 43, close: 49 },
  { name: "BgBlue", open: 44, close: 49 },
  { name: "BgMagenta", open: 45, close: 49 },
  { name: "BgCyan", open: 46, close: 49 },
  { name: "BgWhite", open: 47, close: 49 },
  { name: "BlackBright", open: 90, close: 39 },
  { name: "RedBright", open: 91, close: 39 },
  { name: "GreenBright", open: 92, close: 39 },
  { name: "YellowBright", open: 93, close: 39 },
  { name: "BlueBright", open: 94, close: 39 },
  { name: "MagentaBright", open: 95, close: 39 },
  { name: "CyanBright", open: 96, close: 39 },
  { name: "WhiteBright", open: 97, close: 39 },
  { name: "BgBlackBright", open: 100, close: 49 },
  { name: "BgRedBright", open: 101, close: 49 },
  { name: "BgGreenBright", open: 102, close: 49 },
  { name: "BgYellowBright", open: 103, close: 49 },
  { name: "BgBlueBright", open: 104, close: 49 },
  { name: "BgMagentaBright", open: 105, close: 49 },
  { name: "BgCyanBright", open: 106, close: 49 },
  { name: "BgWhiteBright", open: 107, close: 49 },
];

function r(...args: string[]): string {
  return resolve(dirname(fromFileUrl(import.meta.url)), ...args);
}

function filename(path: string) {
  path = basename(path);
  return path.replace(new RegExp(extname(path)), "");
}

function privateName(name: string) {
  const fw = name[0].toLowerCase();
  return "_" + fw + name.slice(1);
}

async function generateCode(styles: Array<Style>) {
  const FILE_NAME = "style.go";
  for (const it of styles) {
    it.privateName = privateName(it.name);
  }
  const code = await ejs.renderFile(r(`./${filename(FILE_NAME)}.ejs`), {
    styles,
  });
  await Deno.writeFile(r("..", FILE_NAME), new TextEncoder().encode(code));

  const p = Deno.run({
    cmd: ["go", "fmt", FILE_NAME],
    stdin: "inherit",
    stdout: "inherit",
  });
  await p.status();
  console.log(`✨ generate ${FILE_NAME} success!!`);
}

// Learn more at https://deno.land/manual/examples/module_metadata#concepts
if (import.meta.main) {
  generateCode(styles);
}
