const express = require('express');
const passport = require('passport');
const userController = require('../controllers/userController');

const router = express.Router();

router.post('/register', userController.register);
router.post('/login', userController.login);
router.get('/profile', passport.authenticate('jwt', { session: false }), userController.getProfile);

module.exports = router;
