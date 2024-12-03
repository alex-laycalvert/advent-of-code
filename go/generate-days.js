/**
 * This script generates the day files in go for days 1-25 in the days directory.
 */
import { mkdir, writeFile } from "fs/promises";

async function main() {
    const year = process.argv[2];
    const days = Array.from({ length: 25 }, (_, i) => i + 1);
    await mkdir(`./years/year${year}`, { recursive: true });
    await Promise.all(
        days.map((day) =>
            writeFile(`./years/year${year}/day${day}.go`, formatDay(year, day)),
        ),
    );
}

main()
    .then(() => process.exit(0))
    .catch((e) => {
        console.error(e);
        process.exit(1);
    });

function formatDay(year, day) {
    return `package year${year}

type Day${day} struct {
    Input []string
}

func (d Day${day}) Part1() string {
    return "Not Implemented"
}

func (d Day${day}) Part2() string {
    return "Not Implemented"
}
`;
}
