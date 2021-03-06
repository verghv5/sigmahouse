/**
 * Module dependencies.
 */

var template_engine = 'dust',
    domain = 'localhost';

var express = require('express'),
    engine = require('ejs-locals'),
    routes = require('./routes'),
    http = require('http'),
    store = new express.session.MemoryStore,
    path = require('path'),
    hoffman = require('hoffman'),
    requirejs = require('requirejs');


var app = express();

if (template_engine == 'dust') {
    var dust = require('dustjs-linkedin'),
        cons = require('consolidate');

    app.engine('dust', hoffman.__express());
} 


app.configure(function() {

    app.set('template_engine', template_engine);
    app.set('domain', domain);
    app.set('port', process.env.PORT || 8080);
    app.set('views', __dirname + '/views');
    app.set('view engine', template_engine);
    app.set('view cache', true);
    app.use(express.favicon());
    app.use(express.logger('dev'));
    app.use(express.bodyParser());
    app.use(express.methodOverride());
    app.use(express.cookieParser('wigglybits'));
    app.use(express.session({
        secret: 'whatever',
        store: store
    }));
    app.use(app.router);
    app.use(require('less-middleware')(__dirname + '/public'));
 
    app.use(express.static(path.join(__dirname, 'css')));
    app.use(express.static(path.join(__dirname, 'js')));

    hoffman.prime(app.settings.views, function(err) {
        // views are loaded
    });

    //middleware
    app.use(function(req, res, next) {
        if (req.session.user) {
            req.session.logged_in = true;
        }
        res.locals.session = req.session;
        res.locals.q = req.body;
        res.locals.err = false;
        next();
    });

});

app.configure('development', function() {
   app.use(express.errorHandler());
});

app.locals.inspect = require('util').inspect;
app.get('/', routes.index);
app.get('/splash', require('./routes/splash'));

http.createServer(app).listen(app.get('port'), function() {
    console.log("Express server listening on port " + app.get('port'));
});

console.log(routes);
