/**
 * This script generates the day files in go for days 1-25 in the days directory.
 */
import { writeFile } from "fs/promises";

async function main() {
    const days = Array.from({ length: 25 }, (_, i) => i + 1);
    await Promise.all(
        days.map((day) => writeFile(`./days/${day}.go`, formatDay(day))),
    );
}

main()
    .then(() => process.exit(0))
    .catch((e) => {
        console.error(e);
        process.exit(1);
    });

function formatDay(day) {
    return `package days

type Day${day} struct {
    input string
}

func (d Day${day}) Part1() string {
    return "Not Implemented"
}

func (d Day${day}) Part2() string {
    return "Not Implemented"
}
`;
}
