/*
 * Main Display Page
 */

<<<<<<< HEAD
module.exports = function(req, res) {
=======
exports.index = function(req, res) {
>>>>>>> 3160f4f8acd5dbfc1e8920d548044357b2fc70b0
    var template_engine = req.app.settings.template_engine;
    res.locals.session = req.session;
    res.render('splash', {
        title: 'SigmaHouse'
    });
<<<<<<< HEAD
};
=======
};
>>>>>>> 3160f4f8acd5dbfc1e8920d548044357b2fc70b0
