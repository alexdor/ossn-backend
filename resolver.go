//go:generate go run scripts/gqlgen.go -v
package ossn_backend

import (
	"context"

	"github.com/ossn/ossn-backend/models"
)

type Resolver struct{}

func (r *Resolver) Announcement() AnnouncementResolver {
	return &announcementResolver{r}
}
func (r *Resolver) Club() ClubResolver {
	return &clubResolver{r}
}
func (r *Resolver) Event() EventResolver {
	return &eventResolver{r}
}
func (r *Resolver) Job() JobResolver {
	return &jobResolver{r}
}
func (r *Resolver) Location() LocationResolver {
	return &locationResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type announcementResolver struct{ *Resolver }

func (r *announcementResolver) ID(ctx context.Context, obj *models.Announcement) (string, error) {
	return obj.IDToString()
}
func (r *announcementResolver) PublishedAt(ctx context.Context, obj *models.Announcement) (*string, error) {
	return obj.PublishedAtToString()
}
func (r *announcementResolver) CreatedAt(ctx context.Context, obj *models.Announcement) (string, error) {
	return obj.CreatedAtToString(), nil
}
func (r *announcementResolver) UpdatedAt(ctx context.Context, obj *models.Announcement) (string, error) {
	return obj.UpdatedAtToString(), nil
}

type clubResolver struct{ *Resolver }

func (r *clubResolver) ID(ctx context.Context, obj *models.Club) (string, error) {
	return obj.IDToString()
}
func (r *clubResolver) Name(ctx context.Context, obj *models.Club) (*string, error) {
	return obj.Title, nil
}
func (r *clubResolver) Users(ctx context.Context, obj *models.Club) ([]*models.UserWithRole, error) {
	// clubUserRole := []models.ClubUserRole{}
	// usersWithRole := []*models.UserWithRole{}
	// err := models.DBSession.Preload("User").Where("", obj.ID).Find(&clubUserRole).Error
	// if err != nil {
	// 	return clubWithRole, err
	// }

	// for _, club := range clubUserRole {
	// 	c := &club.Club
	// 	usersWithRole = append(usersWithRole, &models.UserWithRole{
	// 		ID:            c.ID.String(),
	// 		Email:         c.Email,
	// 		Location:      c.Location,
	// 		Name:          c.Title,
	// 		ImageURL:      c.ImageURL,
	// 		Role:          &models.Role{Name: models.TurnStringToRolename(club.Role)},
	// 		GithubURL:     c.GithubURL,
	// 		UpdatedAt:     c.UpdatedAtToString(),
	// 		CreatedAt:     c.CreatedAtToString(),
	// 		Description:   c.Description,
	// 		CodeOfConduct: c.CodeOfConduct,
	// 		ClubURL:       c.ClubURL,
	// 		Events:        c.Events,
	// 		// Users:         c.Users,
	// 	})
	// }
	// return usersWithRole, nil
	panic("a")
}
func (r *clubResolver) CreatedAt(ctx context.Context, obj *models.Club) (string, error) {
	return obj.CreatedAtToString(), nil
}
func (r *clubResolver) UpdatedAt(ctx context.Context, obj *models.Club) (string, error) {
	return obj.UpdatedAtToString(), nil
}

type eventResolver struct{ *Resolver }

func (r *eventResolver) ID(ctx context.Context, obj *models.Event) (string, error) {
	return obj.IDToString()
}
func (r *eventResolver) StartDate(ctx context.Context, obj *models.Event) (*string, error) {
	return obj.StartDateToString()
}
func (r *eventResolver) EndDate(ctx context.Context, obj *models.Event) (*string, error) {
	return obj.EndDateToString()
}
func (r *eventResolver) PublishedAt(ctx context.Context, obj *models.Event) (*string, error) {
	return obj.PublishedAtToString()
}
func (r *eventResolver) CreatedAt(ctx context.Context, obj *models.Event) (string, error) {
	return obj.CreatedAtToString(), nil
}
func (r *eventResolver) UpdatedAt(ctx context.Context, obj *models.Event) (string, error) {
	return obj.UpdatedAtToString(), nil
}

type jobResolver struct{ *Resolver }

func (r *jobResolver) ID(ctx context.Context, obj *models.Job) (string, error) {
	return obj.IDToString()
}
func (r *jobResolver) PublishedAt(ctx context.Context, obj *models.Job) (*string, error) {
	return obj.PublishedAtToString()
}
func (r *jobResolver) CreatedAt(ctx context.Context, obj *models.Job) (string, error) {
	return obj.CreatedAtToString(), nil
}
func (r *jobResolver) UpdatedAt(ctx context.Context, obj *models.Job) (string, error) {
	return obj.UpdatedAtToString(), nil
}

type locationResolver struct{ *Resolver }

func (r *locationResolver) ID(ctx context.Context, obj *models.Location) (string, error) {
	return obj.IDToString()
}
func (r *locationResolver) CreatedAt(ctx context.Context, obj *models.Location) (string, error) {
	return obj.CreatedAtToString(), nil
}
func (r *locationResolver) UpdatedAt(ctx context.Context, obj *models.Location) (string, error) {
	return obj.UpdatedAtToString(), nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateClub(ctx context.Context, input models.ClubInput) (*models.Club, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateLocation(ctx context.Context, input *models.LocationInput) (*models.Location, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}
	err := models.DBSession.Where("id = ?", id).First(user).Error
	return user, err
}
func (r *queryResolver) Users(ctx context.Context, limit *int, before *string, after *string, first *int) (*models.Users, error) {
	users := []models.User{}
	err := models.DBSession.Limit(getLimit(limit)).Find(&users).Error
	return &models.Users{Users: users, PageInfo: models.PageInfo{}}, err
}
func (r *queryResolver) Clubs(ctx context.Context, limit *int, userID *string, ids []*string, before *string, after *string, first *int) (*models.Clubs, error) {

	clubs := []models.Club{}
	err := models.DBSession.Limit(getLimit(limit)).Find(&clubs).Error
	if err != nil {
		return nil, err
	}

	return &models.Clubs{Clubs: clubs, PageInfo: models.PageInfo{}}, err
}
func (r *queryResolver) Club(ctx context.Context, id string) (*models.Club, error) {
	club := &models.Club{}
	err := models.DBSession.Where("id =?", id).First(club).Error
	return club, err
}

func (r *queryResolver) Events(ctx context.Context, limit *int, clubId *string, before *string, after *string, first *int) (*models.Events, error) {
	event := []models.Event{}
	err := models.DBSession.Limit(getLimit(limit)).Order("published_at").Find(&event).Error
	if err != nil {
		return nil, err
	}

	return &models.Events{Events: event, PageInfo: models.PageInfo{}}, err
}
func (r *queryResolver) Event(ctx context.Context, id string) (*models.Event, error) {
	event := &models.Event{}
	err := models.DBSession.Where("id =?", id).First(event).Error
	return event, err
}

func (r *queryResolver) Jobs(ctx context.Context, limit *int, before *string, after *string, first *int) (*models.Jobs, error) {

	jobs := []models.Job{}

	err := models.DBSession.Limit(getLimit(limit)).Order("published_at").Find(&jobs).Error
	if err != nil {
		return nil, err
	}

	return &models.Jobs{Jobs: jobs, PageInfo: models.PageInfo{}}, err
}

func (r *queryResolver) Announcements(ctx context.Context, limit *int, before *string, after *string, first *int) (*models.Announcements, error) {
	annanc := []models.Announcement{}
	err := models.DBSession.Limit(getLimit(limit)).Order("published_at").Find(&annanc).Error
	if err != nil {
		return nil, err
	}

	return &models.Announcements{Announcements: annanc, PageInfo: models.PageInfo{}}, err
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return obj.IDToString()
}
func (r *userResolver) Clubs(ctx context.Context, obj *models.User) ([]*models.ClubWithRole, error) {
	clubUserRole := []models.ClubUserRole{}
	clubWithRole := []*models.ClubWithRole{}
	err := models.DBSession.Preload("Club").Preload("Club.Location").Preload("Club.Events").Where("user_id = ?", obj.ID).Find(&clubUserRole).Error
	if err != nil {
		return clubWithRole, err
	}

	for _, club := range clubUserRole {
		c := &club.Club
		users := []*models.User{}
		err := models.DBSession.Raw("SELECT * FROM users where id IN (SELECT user_id from club_user_roles where club_id = ?)", c.ID).Scan(&users).Error
		if err != nil {
			return clubWithRole, err
		}
		clubWithRole = append(clubWithRole, &models.ClubWithRole{
			ID:            c.ID.String(),
			Email:         c.Email,
			Location:      c.Location,
			Name:          c.Title,
			ImageURL:      c.ImageURL,
			Role:          &models.Role{Name: models.TurnStringToRolename(club.Role)},
			GithubURL:     c.GithubURL,
			UpdatedAt:     c.UpdatedAtToString(),
			CreatedAt:     c.CreatedAtToString(),
			Description:   c.Description,
			CodeOfConduct: c.CodeOfConduct,
			ClubURL:       c.ClubURL,
			Events:        c.Events,
			Users:         users,
		})
	}
	return clubWithRole, nil
}
func (r *userResolver) CreatedAt(ctx context.Context, obj *models.User) (string, error) {
	return obj.CreatedAtToString(), nil
}
func (r *userResolver) UpdatedAt(ctx context.Context, obj *models.User) (string, error) {
	return obj.UpdatedAtToString(), nil
}
