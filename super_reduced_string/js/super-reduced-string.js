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

  const dictionary = new Map();

  Array.from(str).forEach(el => {
    if (dictionary.has(el)) {
      dictionary.set(el, dictionary.get(el) + 1);
    } else {
      dictionary.set(el, 1);
    }
  });

  dictionary.forEach((val, key) => {
    if (val % 2 === 0) {
      dictionary.delete(key);
    }
  });

  const res = [...dictionary.keys()].join('');

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
