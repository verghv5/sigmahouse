/**
 * Module dependencies.
 */

var template_engine = 'dust',
    domain = 'localhost';

var express = require('express'),
    engine = require('ejs-locals'),
    routes = require('./routes'),
    http = require('http'),
<<<<<<< HEAD
    session = require('express-session'),
=======
>>>>>>> 3160f4f8acd5dbfc1e8920d548044357b2fc70b0
    store = new express.session.MemoryStore,
    path = require('path'),
    hoffman = require('hoffman'),
    requirejs = require('requirejs');

<<<<<<< HEAD

=======
const bcrypt = require('bcryptjs');
>>>>>>> 3160f4f8acd5dbfc1e8920d548044357b2fc70b0

var app = express();

if (template_engine == 'dust') {
    var dust = require('dustjs-linkedin'),
        cons = require('consolidate');
<<<<<<< HEAD
    app.engine('dust', hoffman.__express());
=======
    app.engine('dust', cons.dust);
>>>>>>> 3160f4f8acd5dbfc1e8920d548044357b2fc70b0

} 


app.configure(function() {

    app.set('template_engine', template_engine);
    app.set('domain', domain);
    app.set('port', process.env.PORT || 8080);
    app.set('views', __dirname + '/views');
    app.set('view engine', template_engine);
    app.set('view cache', true);
<<<<<<< HEAD
=======
    app.engine('dust', hoffman.__express());
>>>>>>> 3160f4f8acd5dbfc1e8920d548044357b2fc70b0
    app.use(express.favicon());
    app.use(express.logger('dev'));
    app.use(express.bodyParser());
    app.use(express.methodOverride());
    app.use(express.cookieParser('wigglybits'));
    app.use(express.session({
        secret: 'whatever',
        store: store
    }));
<<<<<<< HEAD
    app.use(app.router);
    app.use(require('less-middleware')(__dirname + '/public'));
 
    app.use(express.static(path.join(__dirname, 'css')));
    app.use(express.static(path.join(__dirname, 'js')));
=======
    app.use(express.session());
    app.use(app.router);
    app.use(require('less-middleware')(__dirname + '/public'));
    app.use(express.static('public'));
 
    app.use(express.static(path.join(__dirname, 'public')));
>>>>>>> 3160f4f8acd5dbfc1e8920d548044357b2fc70b0

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
<<<<<<< HEAD
   app.use(express.errorHandler());
=======
    app.use(express.errorHandler());
>>>>>>> 3160f4f8acd5dbfc1e8920d548044357b2fc70b0
});

app.locals.inspect = require('util').inspect;
app.get('/', routes.index);
<<<<<<< HEAD
app.get('/splash', require('./routes/splash'));
=======
//app.get('/splash', routes.splash);
>>>>>>> 3160f4f8acd5dbfc1e8920d548044357b2fc70b0

http.createServer(app).listen(app.get('port'), function() {
    console.log("Express server listening on port " + app.get('port'));
});

<<<<<<< HEAD
console.log(routes);
=======
>>>>>>> 3160f4f8acd5dbfc1e8920d548044357b2fc70b0
