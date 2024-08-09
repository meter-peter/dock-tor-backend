const mongoose = require('mongoose');
const bcrypt = require('bcryptjs');
const jwt = require('jsonwebtoken');
const User = require('../models/User');

// Register a new user
exports.register = async (req, res) => {
    const { email, password, fullName, idNumber, role } = req.body;

    try {
        // Check if the user already exists
        const existingUser = await User.findOne({ email });
        if (existingUser) {
            return res.status(400).json({ message: 'User already exists' });
        }

        // Hash the password before saving
        const hashedPassword = await bcrypt.hash(password, 10);

        // Create a new user instance
        const newUser = new User({
            email,
            password: hashedPassword,
            fullName,
            idNumber,
            role
        });

        // Save the user to the database
        await newUser.save();

        res.status(201).json({ message: 'User registered successfully' });
    } catch (error) {
        res.status(500).json({ message: 'Server error' });
    }
};

// Login a user
exports.login = async (req, res) => {
    const { email, password } = req.body;

    try {
        // Find the user by email
        const user = await User.findOne({ email });
        if (!user) {
            return res.status(400).json({ message: 'Invalid credentials' });
        }

        // Compare the provided password with the stored hashed password
        const isMatch = await bcrypt.compare(password, user.password);
        if (!isMatch) {
            return res.status(400).json({ message: 'Invalid credentials' });
        }

        // Create a JWT payload
        const payload = {
            id: user.id,
            role: user.role,
        };

        // Sign the JWT
        const token = jwt.sign(payload, 'your_jwt_secret', { expiresIn: '1h' }); // Replace with your secret key

        // Send the token back to the client
        res.json({ token, role: user.role });
    } catch (error) {
        res.status(500).json({ message: 'Server error' });
    }
};

// Get user profile (JWT protected route)
exports.getProfile = (req, res) => {
    // The user's data is available on req.user thanks to Passport.js
    res.json({
        id: req.user.id,
        email: req.user.email,
        fullName: req.user.fullName,
        idNumber: req.user.idNumber,
        role: req.user.role,
    });
};
