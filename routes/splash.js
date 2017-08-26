/*
 * Main Display Page
 */

exports.index = function(req, res) {
    var template_engine = req.app.settings.template_engine;
    res.locals.session = req.session;
    res.render('splash', {
        title: 'SigmaHouse'
    });
};