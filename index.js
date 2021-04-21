const core = require('@actions/core');
const fs = require('fs');

let modfile = core.getInput('modfile');
if (modfile === null || modfile.length == 0) {
    modfile = "go.mod";
}

const contents = fs.readFileSync(modfile).toString().split("\n");
core.setOutput("modfile", modfile);

const moduleRe = /^module\s+([\w\/\-\.]+)$/g
function extractModule(line) {
    const result = moduleRe.exec(line);
    if (result !== null && result.length >= 2) {
        core.setOutput("module", result[1]);
    }
}

const goVersionRe = /^go\s+([\d\.]+)$/g
function extractGoVersion(line) {
    const result = goVersionRe.exec(line);
    if (result !== null && result.length >= 2) {
        core.setOutput("go_version", result[1]);
    }
}

for (const line of contents) {
    extractModule(line)
    extractGoVersion(line)
}



