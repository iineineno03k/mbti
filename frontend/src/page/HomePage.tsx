import React, { useEffect, useState } from 'react';
import { User } from '../component/User';
import UserTable from '../component/UserTable';
import UserFormModal from '../component/UserFormModal';
import { Button, Typography, Box } from '@mui/material';
import { getMbtiType } from '../component/User';

const HomePage: React.FC = () => {
  const [users, setUsers] = useState<User[]>([]);
  const [currentPage, setCurrentPage] = useState('table'); // 'table', 'mbti'
  const [currentResult, setCurrentResult] = useState(null);
  const [goodUsers, setGoodUsers] = useState<User[]>([]);
  const [badUsers, setBadUsers] = useState<User[]>([]);
  const [showBadUsers, setShowBadUsers] = useState(false);

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const response = await fetch('http://localhost:8080/user');
        const data: User[] = await response.json();
        setUsers(data);
      } catch (error) {
        console.error('Error fetching users:', error);
        setUsers([]);
      }
    };

    fetchUsers();
  }, []);

  const handleCompatibilityCheck = async (user: User) => {
    try {
      const response = await fetch(`http://localhost:8080/user/${user.id}?mbti=${user.mbti}`);
      const result = await response.json();
      setCurrentResult(result);
      setGoodUsers(result.goodUserList || []);
      setBadUsers(result.badUserList || []);
      setCurrentPage('mbti');
      setShowBadUsers(false)
    } catch (error) {
      console.error('Failed to fetch compatibility data:', error);
    }
  };

  return (
    <div>
      <UserFormModal />
      {currentPage === 'table' && <UserTable users={users} onCompatibilityCheck={handleCompatibilityCheck} />}
      {currentPage === 'mbti' && (
        <Box sx={{ textAlign: 'center', mt: 4 }}>
          {currentResult && (
            <>
              <Typography variant="h5" gutterBottom>
                {`${currentResult.user.lastName} ${currentResult.user.firstName}さん (${getMbtiType(currentResult.user.mbti)}) の相性`}
              </Typography>
              <Typography variant="h6" gutterBottom sx={{mt:10}}>
                良い相性一覧
              </Typography>
              <UserTable users={goodUsers} onCompatibilityCheck={handleCompatibilityCheck} />
              <Typography variant="h6" gutterBottom sx={{mt:10}}>
                悪い相性一覧
              </Typography>
              {!showBadUsers ? (
                <Button variant="outlined" color="warning" onClick={() => setShowBadUsers(true)}>
                  閲覧注意。それでも見る
                </Button>
              ) : (
                <UserTable users={badUsers} onCompatibilityCheck={handleCompatibilityCheck} />
              )}
            </>
          )}
          <Box sx={{ justifyContent: 'flex-end', mt: 2}}>
            <Button variant="contained" onClick={() => setCurrentPage('table')}>一覧に戻る</Button>
          </Box>
        </Box>
      )}
    </div>
  );
};

export default HomePage;
