importScripts('./wasm_exec.js')

async function init() {
  const go = new Go()
  const source = await fetch('tinytest.wasm')
  const obj = await WebAssembly.instantiateStreaming(source, go.importObject)
  //const wasm = obj.instance

  self.onmessage = async (event) => {
    if (event.data === 'main') {
      self.postMessage('starting main...')
      go.run(obj.instance)
      self.postMessage('main started')
    }
    if (event.data === 'cancel') {
      self.postMessage('cancelling....')
      obj.instance.exports.cancel()
      self.postMessage('cancelled')
    }
  }
}
init()
