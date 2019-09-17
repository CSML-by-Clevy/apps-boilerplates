const request = require('request-promise');

const myfunction = async function () {
    try {
      const res = await request({
        method: 'GET',
        uri: 'https://randomfox.ca/floof/',
      });
      return JSON.parse(res);
    } catch (error) {
      return { error };
    }
};

exports.handler = async function handler(event) {
    return myfunction();
};
