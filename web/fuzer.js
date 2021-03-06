;(function (factory) {
  if (typeof window !== 'undefined') {
    return factory(window, window)
  } else if (typeof global !== 'undefined') {
    return factory(global)
  } else if (typeof self !== 'undefined') {
    return factory(self)
  } else {
    throw new TypeError('Unknown environment')
  }
})(function (global) {
  let readFile = null

  if (global.require) {
    const fs = require('fs')
    const util = require('util')
    readFile = util.promisify(fs.readFile)
  } else {
    readFile = async function (filePath) {
      return (await fetch(filePath)).arrayBuffer()
    }
  }

  let instance = null

  global.Fuzer = {
    async load(wasmPath) {
      const go = new global.Go()
      const buffer = await readFile(wasmPath)
      const result = await WebAssembly.instantiate(buffer, go.importObject)
      go.run(result.instance)
    },

    getInstance() {
      if (instance) {
        return instance
      }

      instance = global._Fuzer
      delete global._Fuzer
      return instance
    },
  }
})
