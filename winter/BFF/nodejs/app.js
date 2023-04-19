const Koa = require('koa');
const Router = require('koa-router');
const app = new Koa();
const router = new Router();

const client = require('./client');

const sayHello = (data) => {
    return new Promise((resolve, reject) => {
        client.sayHello({name: data}, (err, res) => {
            if (err) reject(err);
            else (resolve(res));
        });
    });
}

const hello = async ctx => {
    const { name } = ctx.params
    let data = await sayHello(name);
    ctx.body = data
}

const about = ctx => {
    ctx.body = 'Hello about'
}

const main = ctx => {
    ctx.body = 'Hello xshop'
}


router.get('/', main);
router.get('/about', about);
router.get('/hello/:name', hello);

app.use(router.routes()).use(router.allowedMethods());

app.listen(3000, ()=>{
    console.log('http://localhost:3000')
})