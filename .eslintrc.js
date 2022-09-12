module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    'plugin:vue/essential',
    '@vue/airbnb',
  ],
  parserOptions: {
    parser: 'babel-eslint',
  },
  rules: {
    //'semi': ['error', 'never'],//分行检测：不需要分号
    //'index': ['warn', 2],//缩进检测：两行缩进，非两行进行警告
    //'quotes': ['error', 'double'],//引号检测：双引号
    'arrow-body-style': ['off'],
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
  },
};
