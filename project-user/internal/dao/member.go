package dao

import (
	"context"
	"gorm.io/gorm"
	"test.com/project-user/internal/data/member"
	"test.com/project-user/internal/database"
	"test.com/project-user/internal/database/gorms"
)

type MemberDAO struct {
	conn *gorms.GormConn
}

func (m *MemberDAO) FindMember(ctx context.Context, account string, password string) (*member.Member, error) {
	var mem *member.Member
	err := m.conn.Session(ctx).Where("account=? and password = ?", account, password).First(&mem).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return mem, err
}

func NewMemberDao() *MemberDAO {
	return &MemberDAO{
		conn: gorms.New(),
	}
}

func (m *MemberDAO) SaveMember(conn database.DbConn, ctx context.Context, mem *member.Member) error {
	m.conn = conn.(*gorms.GormConn)
	return m.conn.Tx(ctx).Create(mem).Error
}

func (m *MemberDAO) GetMemberByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("email=?", email).Count(&count).Error
	//TODO implement me
	return count > 0, err
}

func (m *MemberDAO) GetMemberByAccount(ctx context.Context, account string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("account=?", account).Count(&count).Error
	//TODO implement me
	return count > 0, err
}

func (m *MemberDAO) GetMemberByMobile(ctx context.Context, mobile string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("mobile=?", mobile).Count(&count).Error
	//TODO implement me
	return count > 0, err
}
