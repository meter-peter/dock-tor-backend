const mongoose = require('mongoose');
const Schema = mongoose.Schema;

const medicalHistorySchema = new Schema({
    patient: {
        type: Schema.Types.ObjectId,
        ref: 'Patient',
        required: true
    },
    historyEntries: [{
        type: Schema.Types.ObjectId,
        ref: 'HistoryEntry'
    }]
}, { timestamps: true });

module.exports = mongoose.model('MedicalHistory', medicalHistorySchema);
