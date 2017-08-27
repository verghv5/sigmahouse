/*
 * Main Display Page
 */

module.exports = function(req, res) {
    var issues;
    $.ajax({
        'url': '/issues',
        'type': 'GET',
        'success': function(data) {
            issues = data;
        }
    });
    var template_engine = req.app.settings.template_engine;
    res.locals.session = req.session;
    res.render('splash', issues);
};
