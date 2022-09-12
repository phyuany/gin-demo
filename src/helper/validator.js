const telephoneValidator = (value) => /^1[3|4|5|6|7|8|9]\d{9}$/.test(value);

export default {
  telephoneValidator,
};
