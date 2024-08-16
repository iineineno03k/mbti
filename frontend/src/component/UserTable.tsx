import React from 'react';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Button } from '@mui/material';
import { styled } from '@mui/material/styles';
import { User, getMbtiType } from './User';

const StyledTableContainer = styled(TableContainer)(({ theme }) => ({
  maxWidth: 800,
  margin: 'auto',
  marginTop: theme.spacing(5),
  boxShadow: '0 4px 8px rgba(0,0,0,0.1)'
}));

const StyledTableCell = styled(TableCell)(({ theme }) => ({
  fontSize: 16,
  fontWeight: 'bold',
  color: theme.palette.text.primary
}));

interface UserTableProps {
  users: User[];
}

const UserTable: React.FC<{ users: User[]; onCompatibilityCheck: (user: User) => void }> = ({ users, onCompatibilityCheck }) => {
  return (
    <StyledTableContainer>
      <Table aria-label="simple table">
        <TableHead>
          <TableRow>
            <StyledTableCell>名前</StyledTableCell>
            <StyledTableCell>あだ名</StyledTableCell>
            <StyledTableCell>MBTI</StyledTableCell>
            <StyledTableCell></StyledTableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {users.map((user: User) => (
            <TableRow key={user.id}>
              <StyledTableCell>{user.lastName} {user.firstName}</StyledTableCell>
              <StyledTableCell>{user.nickname}</StyledTableCell>
              <StyledTableCell>{getMbtiType(user.mbti)}</StyledTableCell>
              <StyledTableCell>
                <Button onClick={() => onCompatibilityCheck(user)}>相性を見る</Button>
              </StyledTableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </StyledTableContainer>
  );
};

export default UserTable;
