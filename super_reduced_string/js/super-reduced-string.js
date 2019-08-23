/**
 * Returns super reduced string.
 * @example input: 'aaabccddd' -> output: 'abd'
 * @param {string} str
 * @returns {string}
 */
function superReducedString(str) {
  const empty = 'Empty String';

  if (str.length === 0) {
    return empty;
  }

  const arr = Array.from(str);
  const chars = [];

  arr.forEach(char => {
    if (chars.length === 0) {
      chars.push(char);
    } else if (chars[chars.length - 1] === char) {
      chars.pop();
    } else {
      chars.push(char);
    }
  });

  const res = chars.join('');
  if (res.length === 0) {
    return empty;
  }

  return res;
}

module.exports = superReducedString;

// eslint-disable-next-line no-console
console.log(superReducedString('aaabccddd'));

// eslint-disable-next-line no-console
console.log(superReducedString('aabcc'));
