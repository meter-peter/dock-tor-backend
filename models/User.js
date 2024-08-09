const mongoose = require('mongoose');
const Schema = mongoose.Schema;

const userSchema = new Schema({
    email: {
        type: String,
        required: true,
        unique: true
    },
    password: {
        type: String,
        required: true
    },
    fullName: {
        type: String,
        required: true
    },
    idNumber: {
        type: String,
        required: true,
        unique: true
    },
    role: {
        type: String,
        enum: ['patient', 'secretary', 'doctor'],
        required: true
    }
}, { timestamps: true });

module.exports = mongoose.model('User', userSchema);
