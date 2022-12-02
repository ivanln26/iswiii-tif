import nextJest from "next/jest";
import type { Config } from "jest";

const createJestConfig = nextJest({
  dir: "./",
});

const customJestConfig: Config = {
  moduleDirectories: ["node_modules"],
  moduleNameMapper: {
    "@/(.*)": "<rootDir>/src/$1",
  },
  testEnvironment: "jest-environment-jsdom",
};

export default createJestConfig(customJestConfig);
