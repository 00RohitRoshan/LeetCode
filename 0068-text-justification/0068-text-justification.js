/**
 * @param {string[]} words
 * @param {number} maxWidth
 * @return {string[]}
 */
var fullJustify = function(words, maxWidth) {
    const result = [];
    let line = [];
    let lineLength = 0;

    for (const word of words) {
        if (lineLength + line.length + word.length <= maxWidth) {
            line.push(word);
            lineLength += word.length;
        } else {
            result.push(line);
            line = [word];
            lineLength = word.length;
        }
    }

    result.push(line);

    const justifiedLines = [];
    for (let i = 0; i < result.length - 1; i++) {
        line = result[i];
        const numWords = line.length;
        const numSpaces = maxWidth - line.reduce((acc, word) => acc + word.length, 0);

        let spaceGaps = Math.max(numWords - 1, 1);
        const spacesPerGap = Math.floor(numSpaces / spaceGaps);
        let extraSpaces = numSpaces % spaceGaps;

        let justifiedLine = "";
        for (const word of line) {
            justifiedLine += word;

            if (spaceGaps > 0) {
                const spacesToAdd = spacesPerGap + (extraSpaces > 0 ? 1 : 0);
                justifiedLine += " ".repeat(spacesToAdd);
                extraSpaces -= 1;
                spaceGaps -= 1;
            }
        }

        justifiedLines.push(justifiedLine);
    }

    const lastLine = result[result.length - 1].join(" ");
    justifiedLines.push(lastLine + " ".repeat(maxWidth - lastLine.length));

    return justifiedLines;    
};