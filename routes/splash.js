/*
 * Main Display Page
 */

module.exports = function(req, res) {
    var issues;

    var request = require('request');
    request.get('http://localhost:8000/issues', function(error, response, body) {
        if (error || response.statusCode != 200) {
            console.log('wtf');
            return;
        }
        issues = body;

        var template_engine = req.app.settings.template_engine;
        res.locals.session = req.session;
        res.render('splash', {sigmaIssues: issues});
    });
};

