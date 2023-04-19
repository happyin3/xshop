const { Etcd3 } = require('etcd3');
const client = new Etcd3({hosts: ['127.0.0.1:12379', '127.0.0.1:22379', '127.0.0.1:32379']});

(async () => {
    await client.put('foo').value('bar');

    const fooValue = await client.get('foo').string();
    console.log('foo was:', fooValue);

    const allFValues = await client.getAll().prefix('f').keys();
    console.log('all our keys starting with "f":', allFValues);

    const ss = await client.get('rpc.xshop.hello/127.0.0.1:50051').string();
    console.log('rpc.xshop.hello:', ss);

    await client.delete().all();
})();