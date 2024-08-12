import React, { useState } from 'react';
import { Button, TextField, Modal, Box, Typography, MenuItem, FormControl, InputLabel, Select, SelectChangeEvent } from '@mui/material';
import { User, mbtiTypes } from './User';

const style = {
    position: 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 400,
    bgcolor: 'background.paper',
    border: '2px solid #000',
    boxShadow: 24,
    p: 4,
} as const;


export default function UserFormModal() {
    const [open, setOpen] = useState<boolean>(false);
    const [user, setUser] = useState<User>({ lastName: '', firstName: '', nickname: '', mbti: -1 });

    const handleOpen = () => setOpen(true);
    const handleClose = () => setOpen(false);

    const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        setUser(prev => ({ ...prev, [name]: value }));
    };

    const handleSelectChange = (event: SelectChangeEvent<number>) => {
        const { name, value } = event.target;
        setUser(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmit = async () => {
        if (!user.lastName || !user.firstName || user.mbti === -1) {
            alert('姓、名、MBTIを全て入力してください。');
            return;
        }

        try {
            const response = await fetch('http://localhost:8080/user', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(user)
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            handleClose();
            alert('登録完了しました！');
            window.location.reload();
        } catch (error) {
            alert('Failed to register user: ' + error.message);
        }
    };

    return (
        <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '10vh' }}>
            <Button variant="outlined" onClick={handleOpen}>MBTIを登録する</Button>
            <Modal
                open={open}
                onClose={handleClose}
                aria-labelledby="modal-modal-title"
                aria-describedby="modal-modal-description"
            >
                <Box sx={style}>
                    <Typography id="modal-modal-title" variant="h6" component="h2">
                        あなたの名前とMBTIは？？
                    </Typography>
                    <TextField
                        fullWidth
                        margin="normal"
                        required
                        label="姓"
                        variant="outlined"
                        name="lastName"
                        value={user.lastName}
                        onChange={handleInputChange}
                    />
                    <TextField
                        fullWidth
                        margin="normal"
                        required
                        label="名"
                        variant="outlined"
                        name="firstName"
                        value={user.firstName}
                        onChange={handleInputChange}
                    />
                    <TextField
                        fullWidth
                        margin="normal"
                        label="あだ名"
                        variant="outlined"
                        name="nickname"
                        value={user.nickname}
                        onChange={handleInputChange}
                    />
                    <FormControl fullWidth margin="normal" required>
                        <InputLabel id="mbti-label">あなたのMBTI</InputLabel>
                        <Select
                            labelId="mbti-label"
                            value={user.mbti === -1 ? '' : user.mbti}
                            label="MBTI Type"
                            name="mbti"
                            onChange={handleSelectChange}
                        >
                            {mbtiTypes.map((type, index) => (
                                <MenuItem key={type} value={index}>{type}</MenuItem>
                            ))}
                        </Select>
                    </FormControl>
                    <Box sx={{ display: 'flex', justifyContent: 'flex-end', mt: 2 }}>
                        <Button variant="contained" color="primary" onClick={handleSubmit}>
                            登録
                        </Button>
                    </Box>
                </Box>
            </Modal>
        </div>
    );
}
