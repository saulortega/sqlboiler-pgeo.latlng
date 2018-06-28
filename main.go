// Description:
// https://github.com/saulortega/pgeo
package pgeo

import (
	"database/sql/driver"
	pg "github.com/saulortega/pgeo.latlng"
	"github.com/volatiletech/sqlboiler/randomize"
)

type Point pg.Point

type Circle pg.Circle

type Box pg.Box

type Path pg.Path

type Line pg.Line

type Lseg pg.Lseg

type Polygon pg.Polygon

type NullPoint pg.NullPoint

type NullCircle pg.NullCircle

type NullBox pg.NullBox

type NullPath pg.NullPath

type NullLine pg.NullLine

type NullLseg pg.NullLseg

type NullPolygon pg.NullPolygon

//

func NewPoint(Lat, Lng float64) Point {
	return Point(pg.NewPoint(Lat, Lng))
}

func NewLine(A, B, C float64) Line {
	return Line(pg.NewLine(A, B, C))
}

func NewLseg(A, B Point) Lseg {
	return Lseg(pg.NewLseg(pg.Point(A), pg.Point(B)))
}

func NewBox(A, B Point) Box {
	return Box(pg.NewBox(pg.Point(A), pg.Point(B)))
}

func NewPath(P []Point, C bool) Path {
	return Path(pg.NewPath(convertPoints(P), C))
}

func NewPolygon(P []Point) Polygon {
	return Polygon(pg.NewPolygon(convertPoints(P)))
}

func NewCircle(P Point, R float64) Circle {
	return Circle(pg.NewCircle(pg.Point(P), R))
}

func NewNullPoint(P Point, v bool) NullPoint {
	return NullPoint(pg.NewNullPoint(pg.Point(P), v))
}

func NewNullLine(L Line, v bool) NullLine {
	return NullLine(pg.NewNullLine(pg.Line(L), v))
}

func NewNullLseg(L Lseg, v bool) NullLseg {
	return NullLseg(pg.NewNullLseg(pg.Lseg(L), v))
}

func NewNullBox(B Box, v bool) NullBox {
	return NullBox(pg.NewNullBox(pg.Box(B), v))
}

func NewNullPath(P Path, v bool) NullPath {
	return NullPath(pg.NewNullPath(pg.Path(P), v))
}

func NewNullPolygon(P Polygon, v bool) NullPolygon {
	return NullPolygon(pg.NewNullPolygon(pg.Polygon(P), v))
}

func NewNullCircle(C Circle, v bool) NullCircle {
	return NullCircle(pg.NewNullCircle(pg.Circle(C), v))
}

//
// Point:

func (o Point) Value() (driver.Value, error) {
	return pg.Point(o).Value()
}

func (o *Point) Scan(src interface{}) error {
	val := pg.Point{}
	err := val.Scan(src)
	*o = Point(val)
	return err
}

func (o *Point) Randomize(seed *randomize.Seed, fieldType string, shouldBeNull bool) {
	*o = Point(pg.NewRandPoint())
}

//
// Circle:

func (o Circle) Value() (driver.Value, error) {
	return pg.Circle(o).Value()
}

func (o *Circle) Scan(src interface{}) error {
	val := pg.Circle{}
	err := val.Scan(src)
	*o = Circle(val)
	return err
}

func (o *Circle) Randomize(seed *randomize.Seed, fieldType string, shouldBeNull bool) {
	*o = Circle(pg.NewRandCircle())
}

//
// Box:

func (o Box) Value() (driver.Value, error) {
	return pg.Box(o).Value()
}

func (o *Box) Scan(src interface{}) error {
	val := pg.Box{}
	err := val.Scan(src)
	*o = Box(val)
	return err
}

func (o *Box) Randomize(seed *randomize.Seed, fieldType string, shouldBeNull bool) {
	*o = Box(pg.NewRandBox())
}

//
// Path:

func (o Path) Value() (driver.Value, error) {
	return pg.Path(o).Value()
}

func (o *Path) Scan(src interface{}) error {
	val := pg.Path{}
	err := val.Scan(src)
	*o = Path(val)
	return err
}

func (o *Path) Randomize(seed *randomize.Seed, fieldType string, shouldBeNull bool) {
	*o = Path(pg.NewRandPath())
}

//
// Line:

func (o Line) Value() (driver.Value, error) {
	return pg.Line(o).Value()
}

func (o *Line) Scan(src interface{}) error {
	val := pg.Line{}
	err := val.Scan(src)
	*o = Line(val)
	return err
}

func (o *Line) Randomize(seed *randomize.Seed, fieldType string, shouldBeNull bool) {
	*o = Line(pg.NewRandLine())
}

//
// Lseg:

func (o Lseg) Value() (driver.Value, error) {
	return pg.Lseg(o).Value()
}

func (o *Lseg) Scan(src interface{}) error {
	val := pg.Lseg{}
	err := val.Scan(src)
	*o = Lseg(val)
	return err
}

func (o *Lseg) Randomize(seed *randomize.Seed, fieldType string, shouldBeNull bool) {
	*o = Lseg(pg.NewRandLseg())
}

//
// Polygon:

func (o Polygon) Value() (driver.Value, error) {
	return pg.Polygon(o).Value()
}

func (o *Polygon) Scan(src interface{}) error {
	val := pg.Polygon{}
	err := val.Scan(src)
	*o = Polygon(val)
	return err
}

func (o *Polygon) Randomize(seed *randomize.Seed, fieldType string, shouldBeNull bool) {
	*o = Polygon(pg.NewRandPolygon())
}

//
// NullPoint:

func (o NullPoint) Value() (driver.Value, error) {
	return pg.NullPoint(o).Value()
}

func (o *NullPoint) Scan(src interface{}) error {
	val := pg.NullPoint{}
	err := val.Scan(src)
	*o = NullPoint(val)
	return err
}

func (o *NullPoint) Randomize(seed *randomize.Seed, fieldType string, shouldBeNull bool) {
	*o = NullPoint(pg.NewNullPoint(pg.NewRandPoint(), !shouldBeNull))
}

//
// NullCircle:

func (o NullCircle) Value() (driver.Value, error) {
	return pg.NullCircle(o).Value()
}

func (o *NullCircle) Scan(src interface{}) error {
	val := pg.NullCircle{}
	err := val.Scan(src)
	*o = NullCircle(val)
	return err
}

func (o *NullCircle) Randomize(seed *randomize.Seed, fieldType string, shouldBeNull bool) {
	*o = NullCircle(pg.NewNullCircle(pg.NewRandCircle(), !shouldBeNull))
}

//
// NullBox:

func (o NullBox) Value() (driver.Value, error) {
	return pg.NullBox(o).Value()
}

func (o *NullBox) Scan(src interface{}) error {
	val := pg.NullBox{}
	err := val.Scan(src)
	*o = NullBox(val)
	return err
}

func (o *NullBox) Randomize(seed *randomize.Seed, fieldType string, shouldBeNull bool) {
	*o = NullBox(pg.NewNullBox(pg.NewRandBox(), !shouldBeNull))
}

//
// NullPath:

func (o NullPath) Value() (driver.Value, error) {
	return pg.NullPath(o).Value()
}

func (o *NullPath) Scan(src interface{}) error {
	val := pg.NullPath{}
	err := val.Scan(src)
	*o = NullPath(val)
	return err
}

func (o *NullPath) Randomize(seed *randomize.Seed, fieldType string, shouldBeNull bool) {
	*o = NullPath(pg.NewNullPath(pg.NewRandPath(), !shouldBeNull))
}

//
// NullLine:

func (o NullLine) Value() (driver.Value, error) {
	return pg.NullLine(o).Value()
}

func (o *NullLine) Scan(src interface{}) error {
	val := pg.NullLine{}
	err := val.Scan(src)
	*o = NullLine(val)
	return err
}

func (o *NullLine) Randomize(seed *randomize.Seed, fieldType string, shouldBeNull bool) {
	*o = NullLine(pg.NewNullLine(pg.NewRandLine(), !shouldBeNull))
}

//
// NullLseg:

func (o NullLseg) Value() (driver.Value, error) {
	return pg.NullLseg(o).Value()
}

func (o *NullLseg) Scan(src interface{}) error {
	val := pg.NullLseg{}
	err := val.Scan(src)
	*o = NullLseg(val)
	return err
}

func (o *NullLseg) Randomize(seed *randomize.Seed, fieldType string, shouldBeNull bool) {
	*o = NullLseg(pg.NewNullLseg(pg.NewRandLseg(), !shouldBeNull))
}

//
// NullPolygon:

func (o NullPolygon) Value() (driver.Value, error) {
	return pg.NullPolygon(o).Value()
}

func (o *NullPolygon) Scan(src interface{}) error {
	val := pg.NullPolygon{}
	err := val.Scan(src)
	*o = NullPolygon(val)
	return err
}

func (o *NullPolygon) Randomize(seed *randomize.Seed, fieldType string, shouldBeNull bool) {
	*o = NullPolygon(pg.NewNullPolygon(pg.NewRandPolygon(), !shouldBeNull))
}

//
//
//

func convertPoints(Ps []Point) []pg.Point {
	var points = []pg.Point{}
	if Ps != nil {
		for _, p := range Ps {
			points = append(points, pg.Point(p))
		}
	}

	return points
}
