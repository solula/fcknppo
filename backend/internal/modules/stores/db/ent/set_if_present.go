// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
	"waterfall-backend/internal/models"
	"waterfall-backend/internal/models/files"
	"waterfall-backend/internal/pkg/optional"
)

func (u *ChapterUpdateOne) SetUpdatedAtIfPresent(optionalVar optional.Variable[time.Time]) *ChapterUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetUpdatedAt(*optionalVar.Value())
	}

	return u
}

func (u *ChapterUpdateOne) SetOrClearDeletedAtIfPresent(optionalVar optional.Variable[time.Time]) *ChapterUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetDeletedAt(*optionalVar.Value())
	} else {
		u.ClearDeletedAt()
	}

	return u
}

func (u *ChapterUpdateOne) SetOrClearNumberIfPresent(optionalVar optional.Variable[int]) *ChapterUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetNumber(*optionalVar.Value())
	} else {
		u.ClearNumber()
	}

	return u
}

func (u *ChapterUpdateOne) SetTitleIfPresent(optionalVar optional.Variable[string]) *ChapterUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetTitle(*optionalVar.Value())
	}

	return u
}

func (u *ChapterUpdateOne) SetPartUUIDIfPresent(optionalVar optional.Variable[string]) *ChapterUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetPartUUID(*optionalVar.Value())
	}

	return u
}

func (u *ChapterUpdateOne) SetOrClearReleaseUUIDIfPresent(optionalVar optional.Variable[string]) *ChapterUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetReleaseUUID(*optionalVar.Value())
	} else {
		u.ClearReleaseUUID()
	}

	return u
}

func (u *CommentUpdateOne) SetUpdatedAtIfPresent(optionalVar optional.Variable[time.Time]) *CommentUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetUpdatedAt(*optionalVar.Value())
	}

	return u
}

func (u *CommentUpdateOne) SetOrClearDeletedAtIfPresent(optionalVar optional.Variable[time.Time]) *CommentUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetDeletedAt(*optionalVar.Value())
	} else {
		u.ClearDeletedAt()
	}

	return u
}

func (u *CommentUpdateOne) SetTextIfPresent(optionalVar optional.Variable[string]) *CommentUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetText(*optionalVar.Value())
	}

	return u
}

func (u *CommentUpdateOne) SetAuthorUUIDIfPresent(optionalVar optional.Variable[string]) *CommentUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetAuthorUUID(*optionalVar.Value())
	}

	return u
}

func (u *CommentUpdateOne) SetOrClearParentUUIDIfPresent(optionalVar optional.Variable[string]) *CommentUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetParentUUID(*optionalVar.Value())
	} else {
		u.ClearParentUUID()
	}

	return u
}

func (u *CommentUpdateOne) SetOrClearChapterUUIDIfPresent(optionalVar optional.Variable[string]) *CommentUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetChapterUUID(*optionalVar.Value())
	} else {
		u.ClearChapterUUID()
	}

	return u
}

func (u *FileUpdateOne) SetUpdatedAtIfPresent(optionalVar optional.Variable[time.Time]) *FileUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetUpdatedAt(*optionalVar.Value())
	}

	return u
}

func (u *FileUpdateOne) SetFilenameIfPresent(optionalVar optional.Variable[string]) *FileUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetFilename(*optionalVar.Value())
	}

	return u
}

func (u *FileUpdateOne) SetMimeTypeIfPresent(optionalVar optional.Variable[string]) *FileUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetMimeType(*optionalVar.Value())
	}

	return u
}

func (u *FileUpdateOne) SetDescriptionIfPresent(optionalVar optional.Variable[string]) *FileUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetDescription(*optionalVar.Value())
	}

	return u
}

func (u *FileUpdateOne) SetOrClearCreatorUUIDIfPresent(optionalVar optional.Variable[string]) *FileUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetCreatorUUID(*optionalVar.Value())
	} else {
		u.ClearCreatorUUID()
	}

	return u
}

func (u *FileUpdateOne) SetObjectTypeIfPresent(optionalVar optional.Variable[models.ObjectType]) *FileUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetObjectType(*optionalVar.Value())
	}

	return u
}

func (u *FileUpdateOne) SetObjectRefIfPresent(optionalVar optional.Variable[string]) *FileUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetObjectRef(*optionalVar.Value())
	}

	return u
}

func (u *FileUpdateOne) SetTypeIfPresent(optionalVar optional.Variable[files.Type]) *FileUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetType(*optionalVar.Value())
	}

	return u
}

func (u *FileUpdateOne) SetTempIfPresent(optionalVar optional.Variable[bool]) *FileUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetTemp(*optionalVar.Value())
	}

	return u
}

func (u *FileUpdateOne) SetOrClearSequenceNumberIfPresent(optionalVar optional.Variable[uint]) *FileUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetSequenceNumber(*optionalVar.Value())
	} else {
		u.ClearSequenceNumber()
	}

	return u
}

func (u *MigrationsUpdateOne) SetMigratedIfPresent(optionalVar optional.Variable[int]) *MigrationsUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetMigrated(*optionalVar.Value())
	}

	return u
}

func (u *PartUpdateOne) SetUpdatedAtIfPresent(optionalVar optional.Variable[time.Time]) *PartUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetUpdatedAt(*optionalVar.Value())
	}

	return u
}

func (u *PartUpdateOne) SetOrClearDeletedAtIfPresent(optionalVar optional.Variable[time.Time]) *PartUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetDeletedAt(*optionalVar.Value())
	} else {
		u.ClearDeletedAt()
	}

	return u
}

func (u *PartUpdateOne) SetOrClearNumberIfPresent(optionalVar optional.Variable[int]) *PartUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetNumber(*optionalVar.Value())
	} else {
		u.ClearNumber()
	}

	return u
}

func (u *PartUpdateOne) SetTitleIfPresent(optionalVar optional.Variable[string]) *PartUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetTitle(*optionalVar.Value())
	}

	return u
}

func (u *PartUpdateOne) SetOrClearAnnotationIfPresent(optionalVar optional.Variable[string]) *PartUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetAnnotation(*optionalVar.Value())
	} else {
		u.ClearAnnotation()
	}

	return u
}

func (u *PartUpdateOne) SetOrClearReleaseUUIDIfPresent(optionalVar optional.Variable[string]) *PartUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetReleaseUUID(*optionalVar.Value())
	} else {
		u.ClearReleaseUUID()
	}

	return u
}

func (u *PermissionUpdateOne) SetDescriptionIfPresent(optionalVar optional.Variable[string]) *PermissionUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetDescription(*optionalVar.Value())
	}

	return u
}

func (u *ReleaseUpdateOne) SetUpdatedAtIfPresent(optionalVar optional.Variable[time.Time]) *ReleaseUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetUpdatedAt(*optionalVar.Value())
	}

	return u
}

func (u *ReleaseUpdateOne) SetReleaseDateIfPresent(optionalVar optional.Variable[time.Time]) *ReleaseUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetReleaseDate(*optionalVar.Value())
	}

	return u
}

func (u *ReleaseUpdateOne) SetDescriptionIfPresent(optionalVar optional.Variable[string]) *ReleaseUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetDescription(*optionalVar.Value())
	}

	return u
}

func (u *RoleUpdateOne) SetDescriptionIfPresent(optionalVar optional.Variable[string]) *RoleUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetDescription(*optionalVar.Value())
	}

	return u
}

func (u *RoleUpdateOne) SetReleaseDelayIfPresent(optionalVar optional.Variable[float64]) *RoleUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetReleaseDelay(*optionalVar.Value())
	}

	return u
}

func (u *SeedMigrationsUpdateOne) SetMigratedIfPresent(optionalVar optional.Variable[int]) *SeedMigrationsUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetMigrated(*optionalVar.Value())
	}

	return u
}

func (u *UserUpdateOne) SetUpdatedAtIfPresent(optionalVar optional.Variable[time.Time]) *UserUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetUpdatedAt(*optionalVar.Value())
	}

	return u
}

func (u *UserUpdateOne) SetOrClearDeletedAtIfPresent(optionalVar optional.Variable[time.Time]) *UserUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetDeletedAt(*optionalVar.Value())
	} else {
		u.ClearDeletedAt()
	}

	return u
}

func (u *UserUpdateOne) SetOrClearEmailIfPresent(optionalVar optional.Variable[string]) *UserUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetEmail(*optionalVar.Value())
	} else {
		u.ClearEmail()
	}

	return u
}

func (u *UserUpdateOne) SetFullnameIfPresent(optionalVar optional.Variable[string]) *UserUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetFullname(*optionalVar.Value())
	}

	return u
}

func (u *UserUpdateOne) SetUsernameIfPresent(optionalVar optional.Variable[string]) *UserUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetUsername(*optionalVar.Value())
	}

	return u
}

func (u *UserUpdateOne) SetOrClearPasswordHashIfPresent(optionalVar optional.Variable[string]) *UserUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetPasswordHash(*optionalVar.Value())
	} else {
		u.ClearPasswordHash()
	}

	return u
}

func (u *UserUpdateOne) SetOrClearVkIDIfPresent(optionalVar optional.Variable[int64]) *UserUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetVkID(*optionalVar.Value())
	} else {
		u.ClearVkID()
	}

	return u
}

func (u *UserUpdateOne) SetScoreIfPresent(optionalVar optional.Variable[int]) *UserUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetScore(*optionalVar.Value())
	}

	return u
}

func (u *UserUpdateOne) SetEmailVerifiedIfPresent(optionalVar optional.Variable[bool]) *UserUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetEmailVerified(*optionalVar.Value())
	}

	return u
}

func (u *UserUpdateOne) SetSerialNumberIfPresent(optionalVar optional.Variable[uint]) *UserUpdateOne {
	if !optionalVar.IsSet() {
		return u
	}

	if optionalVar.Value() != nil {
		u.SetSerialNumber(*optionalVar.Value())
	}

	return u
}
