const superReducedString = require('./super-reduced-string');

const testCases = [
  {
    input: 'aabcc',
    output: 'b',
    name: 'aabcc',
  },
  {
    input: 'aacc',
    output: 'Empty String',
    name: 'aacc',
  },
  {
    input: '',
    output: 'Empty String',
    name: 'An empty string',
  },
  {
    input: 'aaabccddd',
    output: 'abd',
    name: 'aaabccddd',
  },
  {
    // second try to eliminate flaky case
    input: 'aaabccddd',
    output: 'abd',
    name: 'aaabccddd',
  },
  {
    input: 'aa',
    output: 'Empty String',
    name: 'aa',
  },
  {
    input: 'baab',
    output: 'Empty String',
    name: 'baab',
  },
  {
    input: 'aaaaabbccddd',
    output: 'ad',
    name: 'aaaaabbccddd',
  },
];

describe('superReducedString function should return the right string', () => {
  testCases.forEach(testCase => {
    const { input, output, name } = testCase;
    test(`${name} => ${output}`, () => {
      expect(superReducedString(input)).toBe(output);
    });
  });
});
