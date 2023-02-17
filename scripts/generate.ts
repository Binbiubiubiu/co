import path from 'node:path'
import { fileURLToPath } from 'node:url';
import { Color } from "./types.ts";

const colors: Array<Color> = [
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
  { name: "RedBrCyanBrightight", open: 96, close: 39 },
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

function r(...args:string[]):string{
  return path.resolve(path.dirname(fileURLToPath(import.meta.url)),...args)
}

async function generateCode(colors: Array<Color>) {
  let code = `
  package co\n\n
type Colors struct {\n
  `;
  for (const co of colors) {
    code += `${co.name}           ColorFunc\n`;
  }
  code += `
}\n
func noop(s string) string {\n
	return s\n
}\n
\n
var noColors = Colors{\n
`;
  for (const co of colors) {
    code += `${co.name}:noop,\n`;
  }
  code += `

}\n

var colors = Colors{\n
`;
  for (const co of colors) {
    code += `${co.name},\n`;
  }
  code += `}\n`;

  for (const co of colors) {
    code += `var ${co.name} = ${
      co.replace
        ? `buildWithReplace(${co.open}, ${co.close},${co.replace})`
        : `build(${co.open}, ${co.close})`
    }\n`;
  }

  const encoder = new TextEncoder();
  const data = encoder.encode(code);
  Deno.cwd()
  await Deno.writeFile(r("../color.go"),data)

  Deno.run({
    cmd:["go","fmt","color.go"]
  })

}

// Learn more at https://deno.land/manual/examples/module_metadata#concepts
if (import.meta.main) {
  generateCode(colors);
}
